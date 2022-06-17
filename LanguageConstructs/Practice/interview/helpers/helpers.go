package helpers

import "fmt"


// the length of buffer needs to be at least the
// length of the src slice
func CopySlice(slice interface{}) (data []int, err error) {
   defer func(){
    if p := recover(); p != nil {
      err = fmt.Errorf("recovered!")
   }
   }()

   s, ok := slice.([]int)
   if !ok {
     panic("expected slice of type []int")
   }

   var buffer = make([]int, len(s))

   copy(buffer, s)

   return buffer, nil
}
