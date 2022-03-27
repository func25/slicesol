package boardsol

//DFS returns selected elements that meet specific condition
//if keepExists == true and the pos was already selected then BFS will return nil slice
func (q *SearchQuery[T]) DFS(pos Vector2D, opts ...BFSOption[T]) ([]Vector2D, error) {
	// apply options
	for i := range opts {
		opts[i](q)
	}

	// check if the pos is already iterated before
	// if true, then return a nil slice
	if q.keepExists && q.exists != nil {
		if _, exist := q.exists[pos.to1D(q.Width)]; exist {
			return nil, nil
		}
	}

	if err := q.validate(pos); err != nil {
		return nil, err
	}

	q.internalDFS(pos)
	return q.res, nil
}

func (q *SearchQuery[T]) internalDFS(curPos Vector2D) bool {
	if q.limit != 0 && q.limit <= len(q.res) {
		return false
	}

	dirs := q.directions
	if q.genDirections != nil {
		dirs = q.genDirections()
	}

	for _, dir := range dirs {
		newPos := curPos.plus(dir)
		id1D := newPos.to1D(q.Width)
		if _, exist := q.exists[id1D]; !exist && q.SelectCondition(curPos, newPos) {
			q.res = append(q.res, newPos)
			q.exists[id1D] = emptyStruct

			if !q.internalDFS(newPos) {
				return false
			}
		}
	}

	return true
}
