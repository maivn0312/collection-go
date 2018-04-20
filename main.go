package collections

type Collection map[int]interface{}

func New() Collection {
	collection := Collection{}
	return collection
}

func (collection *Collection) Push(elements... interface{}) {
	for _, element := range elements {
		(*collection)[len(*collection)] = element
	}
}

func (collection *Collection) Remove(elements... interface{}) {
	for i, tmpElement := range *collection {
		for _, element := range elements {
			if tmpElement == element {
				delete(*collection, i)
			}
		}
	}
}

func (collection *Collection) Filter(callback func(entity interface{}) bool) Collection {
	resColleciton := Collection{}
	for _, tmpEntity := range *collection {
		if callback(tmpEntity) == true {
			resColleciton.Push(tmpEntity)
		}
	}
	return resColleciton
}

