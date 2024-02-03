package store

import (
	"testing"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

func TestName_MarshalBSONValue(t *testing.T) {
	in := "El Corral"
	name := Name{Value: in, Valid: true}

	gotTyp, gotVal, err := name.MarshalBSONValue()
	if err != nil {
		t.Error(err)
	}

	wantTyp, wantVal, _ := bson.MarshalValue(in)

	if err != nil || gotTyp != wantTyp || !cmp.Equal(gotVal, wantVal) {
		t.Errorf("store.Name.MarshalBSONValue() = %s %s %v, want %s %s nil", gotTyp, gotVal, err, wantTyp, wantVal)
	}
}

func TestName_UnmarshalBSONValue(t *testing.T) {
	want := Name{Value: "Burger King", Valid: true}
	typ, data, err := bson.MarshalValue(want)

	got := Name{}

	err = bson.UnmarshalValue(typ, data, &got)

	if err != nil || !cmp.Equal(got, want) {
		t.Errorf("store.Name.UnmarshalBSONValue() = %v %v, want %v nil\n%v", got, err, want, cmp.Diff(want, got))
	}
}
