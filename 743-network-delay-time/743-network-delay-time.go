type Distance struct {
    W int
    V int
}

type DistanceHeap []Distance

func (h DistanceHeap) Len() int { return len(h) }
func (h DistanceHeap) Less(i, j int) bool { return h[i].W < h[j].W }
func (h DistanceHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *DistanceHeap) Push(x interface{}) {
    *h = append(*h, x.(Distance))
}

func (h *DistanceHeap) Pop() interface{} {
    value := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return value
}

func networkDelayTime(times [][]int, N int, K int) int {
    graph := map[int]DistanceHeap{}
    
    // keep data node in graph connected other node anything
    // graph[2] = [{ W: 1, V: 3 }, {W: 1, V: 1}]  -> node 2 connected node 3 weight 1 and node 1 weight 1
    for _, time := range times {
        u, v, w := time[0], time[1], time[2]
        graph[u] = append(graph[u], Distance{w, v})
    }
    
    // initial distance start node 2 weight0
    q := &DistanceHeap{
        Distance{0, K},
    }
    
    arrivalTime := map[int]int{}
    
    // bfs process
    for q.Len() > 0 {
        // pop bfs queue
        d := heap.Pop(q).(Distance)

        if _, ok := arrivalTime[d.V]; ok {
            continue
        }
        
        // update time to arrival
        arrivalTime[d.V] = d.W
        
        // update next queue and accumulate time
        for _, next := range graph[d.V] {
            next.W += d.W
            heap.Push(q, next)
        }
    }
    
    // visit not complete all nodes
    if len(arrivalTime) < N {
        return -1
    }
    
    maxTime := 0
    for _, v := range arrivalTime {
        if v > maxTime {
            maxTime = v
        }
    }
    
    return maxTime
}