package oss

import "time"


type Object struct {
	/*object name starting with '/'*/
	Name string

	Size uint64

	/*start with '/', without bucket part */
	ParentDir string

	IsDir bool

	/*modify time */
	ModifyTime time.Time

	MD5 string
}


func ListObject(client *Client, bucket string) (objects []Object, err error) {

}
