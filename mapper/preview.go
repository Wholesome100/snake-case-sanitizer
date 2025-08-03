package mapper

import (
	"sync"
)

// Concurrency-safe Map to store original and modified filenames
var nameMap sync.Map

// TODO: implement a preview feature for the user so they can see changes before application
// Changes applied must be done atomically, if there is a failure at any point then the original names are rolled back
