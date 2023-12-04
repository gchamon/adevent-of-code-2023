package utils

import "testing"

func TestAddSetElements(t *testing.T) {
	intToAdd := []int{1, 2, 3}

	t.Run("int test one by one", func(t *testing.T) {
		setInt := NewSet[int]()
		for _, elm := range intToAdd {
			setInt.Add(elm)
			if !setInt.Exists(elm) {
				t.Errorf("should have been able to add %d to set %+v", elm, setInt)
			}
		}
	})
	t.Run("int test one by one", func(t *testing.T) {
		setInt := NewSet[int]()
		setInt.Add(intToAdd...)
		for _, elm := range intToAdd {
			if !setInt.Exists(elm) {
				t.Errorf("should have been able to add %d to set %+v", elm, setInt)
			}
		}
	})
}

func TestRemoveElements(t *testing.T) {

}
