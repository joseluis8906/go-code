-- Read the docs: https://www.lunarvim.org/docs/configuration
-- Video Tutorials: https://www.youtube.com/watch?v=sFA9kX-Ud_c&list=PLhoH5vyxr6QqGu0i7tt_XoVK9v-KvZ3m6
-- Forum: https://www.reddit.com/r/lunarvim/
-- Discord: https://discord.com/invite/Xb9B4Ny
--

vim.opt.relativenumber = true
vim.opt.foldmethod = "expr"
vim.opt.foldexpr = "nvim_treesitter#foldexpr()"
vim.opt.cursorline = true
vim.opt.cursorcolumn = true
vim.opt.colorcolumn = "120"

lvim.colorscheme = "lunar"
lvim.builtin.nvimtree.setup.view.side = "right"
lvim.builtin.nvimtree.setup.renderer.icons.show.git = false
lvim.format_on_save = true
lvim.keys.normal_mode["<S-E>"] = ":Explore<CR>"

lvim.plugins = {
  { "mxsdev/nvim-dap-vscode-js" },
  { "leoluz/nvim-dap-go" },
  { "hrsh7th/nvim-cmp" },
  { "hrsh7th/cmp-nvim-lsp" },
  -- Copilot plugins are defined below:
  {
    "zbirenbaum/copilot.lua",
    cmd = "Copilot",
    event = "InsertEnter",
    config = function()
      require("copilot").setup({})
    end,
  },
  {
    "zbirenbaum/copilot-cmp",
    config = function()
      require("copilot_cmp").setup({
        suggestion = { enabled = true },
        panel = { enabled = true }
      })
    end
  },
  {
    "chrishrb/gx.nvim",
    event = { "BufEnter" },
    dependencies = { "nvim-lua/plenary.nvim" },
    config = true, -- default settings
  },
}

-- Below config is required to prevent copilot overriding Tab with a suggestion
-- when you're just trying to indent!
local has_words_before = function()
  if vim.api.nvim_buf_get_option(0, "buftype") == "prompt" then return false end
  local line, col = unpack(vim.api.nvim_win_get_cursor(0))
  return col ~= 0 and vim.api.nvim_buf_get_text(0, line - 1, 0, line - 1, col, {})[1]:match("^%s*$") == nil
end

local on_tab = vim.schedule_wrap(function(fallback)
  local cmp = require("cmp")
  if cmp.visible() and has_words_before() then
    cmp.select_next_item({ behavior = cmp.SelectBehavior.Select })
  else
    fallback()
  end
end)

lvim.builtin.cmp.mapping["<Tab>"] = on_tab
-- end copilot conf

-- formaters
local formatters = require "lvim.lsp.null-ls.formatters"
formatters.setup {
  {
    name = "goimports",
    filetypes = { "go" },
  }
}
-- end formaters

-- linters
local linters = require "lvim.lsp.null-ls.linters"
linters.setup {
  {
    command = "golangci-lint",
    filetypes = { "go" },
  },
}
-- end linters

local lsp_capabilities = require("cmp_nvim_lsp").default_capabilities()
vim.list_extend(lvim.lsp.automatic_configuration.skipped_servers, { "gopls" })
local opts = {
  cmd = { "gopls", "-remote=auto" },
  flags = {
    debounce_text_changes = 1000,
  },
  init_options = {
    staticcheck = true,
    gofumpt = true,
    hoverKind = "FullDocumentation",
    -- env = {
    --   GOPACKAGESDRIVER = './tools/gopackagesdriver.sh',
    -- },
    -- directoryFilters = {
    --   "-bazel-bin",
    --   "-bazel-out",
    --   "-bazel-testlogs",
    --   "-bazel-mypkg",
    -- },
  },
  capabilities = lsp_capabilities,
}

require("lvim.lsp.manager").setup("gopls", opts)

local cmp = require('cmp')
cmp.setup {
  sources = cmp.config.sources({
    { name = "nvim_lsp" },
  }),
}

local mason_path = vim.fn.glob(vim.fn.stdpath "data" .. "/mason/")
require("dap-vscode-js").setup {
  debugger_path = mason_path .. "packages/js-debug-adapter",                                   -- Path to vscode-js-debug installation.
  adapters = { "pwa-node", "pwa-chrome", "pwa-msedge", "node-terminal", "pwa-extensionHost" }, -- which adapters to register in nvim-dap
}

for _, language in ipairs { "typescript", "javascript" } do
  require("dap").configurations[language] = {
    {
      type = "pwa-node",
      request = "launch",
      name = "Launch file",
      program = "${file}",
      cwd = "${workspaceFolder}",
    },
    {
      type = "pwa-node",
      request = "attach",
      name = "Attach",
      processId = require 'dap.utils'.pick_process,
      cwd = "${workspaceFolder}",
    },
    {
      type = "pwa-node",
      request = "launch",
      name = "Debug Jest Tests",
      runtimeExecutable = "node",
      runtimeArgs = {
        "./node_modules/jest/bin/jest.js",
        "--runInBand",
      },
      rootPath = "${workspaceFolder}",
      cwd = "${workspaceFolder}",
      console = "integratedTerminal",
      internalConsoleOptions = "neverOpen",
    },
  }
end

local dap = require("dap")

dap.adapters.go = {
  type = 'server',
  port = 2345
}

dap.configurations.go = {
  {
    type = "go",
    name = "Attach Remote",
    request = "attach",
    mode = "remote",
    port = "2345",
  },
}
