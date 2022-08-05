package main

// visited是在入队之前进行判断，确保重复的不进队列，更简洁
// 入队之后就对visited进行修改
func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]int, 0)
	for _, v := range deadends {
		deadendsMap[v] = 1
	}
	visited := make(map[string]int, 0)
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited["0000"] = 1
	var item string
	step := 1
	for ;len(queue) > 0;{
		qSize := len(queue)
		for i:= 0;i<qSize;i++ {
			item, queue = queue[0], queue[1:]
			if item == target {
				return step
			}
			if _, ok := deadendsMap[item]; ok {
				continue
			}
			// item相邻元素入队
			enqueueNeighbors(item, &queue, visited)
		}
		step += 1
	}
	return -1
}


func enqueueNeighbors(item string, queue *[]string, visited map[string]int) {
	for i:=0;i<len(item);i++ {
		tmpMinus := minus(item, i)
		tmpAdd := add(item, i)
		if _, ok := visited[tmpMinus]; !ok {
			*queue = append(*queue, tmpMinus)
			visited[tmpMinus] = 1
		}
		if _, ok := visited[tmpAdd]; !ok{
			*queue = append(*queue, tmpAdd)
			visited[tmpAdd] = 1
		}
	}
}

func minus(item string, index int) string {
	ch := item[index]
	if ch == '0' {
		ch = '9'
	} else {
		ch -= 1
	}
	return item[:index] + string(ch) + item[index+1:]
}

func add(item string, index int) string {
	ch := item[index]
	if ch == '9' {
		ch = '0'
	} else {
		ch += 1
	}
	return item[:index] + string(ch) + item[index+1:]
}

func main() {
	deadends := []string{"0201","0101","0102","1212","2002"}
	openLock(deadends, "0202")
}
