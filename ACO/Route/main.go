package main

import (
	"fmt"
	"math/rand"
	"math"
	// "time"
)

const AntCount = 10 // 蚂蚁的个数
const Alpha = 1 //信息素重要程度参数
const Beta =5 //启发式因子重要程度参数
const Rho =0.1 //信息素蒸发系数
// const G = 200 //最大迭代次数
const Q = 100 // 信息素增强强度系数
// const city = 5 //城市的个数
const MaxIterations = 100
const MaxNum =4

type Ant struct {
	ID         int
	Visited    []bool
	Tour       []int
	TotalCost  float64
	CurrentPos int
	StartPos int
	EndPos int
}

func NewAnt(id, size int) *Ant {
	visited := make([]bool, size)
	tour := make([]int, 0)
	// currentPos := rand.Intn(size)
	visited[0] = true
	tour = append(tour, 0)
	return &Ant{
		ID:         id,
		Visited:    visited,
		Tour:       tour,
		TotalCost:  0.0,
		CurrentPos: 0,
		StartPos: 0,
		EndPos: 4,
	}
}

func (ant *Ant) MoveTo(city int, distanceMatrix [][]float64) {
	if city == -1 {
		// fmt.Printf("选路有问题	")
		return
	}else {
		
		ant.Tour = append(ant.Tour, city)
		ant.TotalCost += distanceMatrix[ant.CurrentPos][city]
		ant.CurrentPos = city
		ant.Visited[city] = true
		
	}
}

func (ant *Ant) CalculateTotalCost(distanceMatrix [][]float64) {
	ant.TotalCost = 0.0
	for i := 0; i < len(ant.Tour)-1; i++ {
		from := ant.Tour[i]
		to := ant.Tour[i+1]
		ant.TotalCost += distanceMatrix[from][to]
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	size := 5 // 图中节点的数量

	// 初始化距离矩阵
	distanceMatrix := make([][]float64, size)
	for i := range distanceMatrix {
		distanceMatrix[i] = make([]float64, size)
		for j := range distanceMatrix[i] {
			distanceMatrix[i][j] = math.Inf(1)
		}
	}
	// 在这里填充 distanceMatrix，表示节点之间的距离
	distanceMatrix[0][0] = 0
	distanceMatrix[0][1] = 9    // A to B
	distanceMatrix[0][2] = 5    // A to C
	distanceMatrix[0][3] = 4    // A to D
	distanceMatrix[0][4] = math.Inf(1)
	distanceMatrix[1][0] = 9
	distanceMatrix[1][1] = 0
	distanceMatrix[1][2] = 2    // B to C
	distanceMatrix[1][3] = math.Inf(1) // B to D (not reachable)
	distanceMatrix[1][4] = 8    // B to E
	distanceMatrix[2][0] = 5    // C to A
	distanceMatrix[2][1] = 2    // C to B
	distanceMatrix[2][2] = 0
	distanceMatrix[2][3] = 3    // C to D
	distanceMatrix[2][4] = 7    // C to E
	distanceMatrix[3][0] = 4    // D to A
	distanceMatrix[3][1] = math.Inf(1) // D to B (not reachable)
	distanceMatrix[3][2] = 3    // D to C
	distanceMatrix[3][3] = 0
	distanceMatrix[3][4] = 1    // D to E
	distanceMatrix[4][0] = math.Inf(1) // E to A (not reachable)
	distanceMatrix[4][1] = 8    // E to B
	distanceMatrix[4][2] = 7    // E to C
	distanceMatrix[4][3] = 1    // E to D
	distanceMatrix[4][4] = 0


	// 初始化信息素浓度
	pheromoneMatrix := make([][]float64, size)
	for i := range pheromoneMatrix {
		pheromoneMatrix[i] = make([]float64, size)
		for j := range pheromoneMatrix[i] {
			pheromoneMatrix[i][j] = 1.0 // 可以将初始信息素浓度设为1.0或其他适当值
		}
	}

	// 主循环
	for iteration := 0; iteration < MaxIterations; iteration++ {
		ants := make([]*Ant, AntCount)

		for i := 0; i < AntCount; i++ {
			ants[i] = NewAnt(i, size)
		}

		// for len(ants[0].Tour) < size+1 {
		// 	for _, ant := range ants {
		// 		ant.MoveTo(SelectNextCity(ant, pheromoneMatrix, distanceMatrix), distanceMatrix)
		// 	}
		// }

		for i:=0; i<MaxNum; i++ {
			for _, ant := range ants {
				if ant.CurrentPos == ant.EndPos{
					continue
				}else{
					ant.MoveTo(SelectNextCity(ant, pheromoneMatrix, distanceMatrix), distanceMatrix)
				}
			}
		}
	

		UpdatePheromone(pheromoneMatrix,  distanceMatrix, ants)

		bestAnt := findBestAnt(ants)
		fmt.Printf("Iteration %d: Best Tour Length: %.2f\n", iteration, bestAnt.TotalCost)
		for _, point := range bestAnt.Tour {
			fmt.Printf("%d ",point)
		}
		// 重新初始化蚂蚁
		for _, ant := range ants {
			ant.Visited = make([]bool, size)
			ant.Tour = []int{}
			ant.TotalCost = 0.0
			ant.CurrentPos = rand.Intn(size)
		}
	}
}

func SelectNextCity(ant *Ant, pheromoneMatrix [][]float64, distanceMatrix [][]float64) int {
	// 当前城市
	currentCity := ant.CurrentPos

	// 未访问的城市列表
	unvisitedCities := make([]int, 0)
	for city := 0; city < len(ant.Visited); city++ {
		if !ant.Visited[city] {
			unvisitedCities = append(unvisitedCities, city)
		}
	}
	// 选择下一个城市
	selectedCity := -1
	// if len(unvisitedCities)==0{
	// 	// fmt.Printf("%d    ",ant.StartPos)
	// 	// fmt.Printf("%d current cities   ",currentCity)
	// 	// fmt.Printf("%f distance ",distanceMatrix[currentCity][ant.StartPos])
	// 	if distanceMatrix[currentCity][ant.StartPos] == math.Inf(1){
	// 		ant.Tour=nil
	// 		ant.TotalCost = math.Inf(1)
	// 		return selectedCity
	// 	} else {
	// 		return ant.StartPos

	// 	}
		
	// }else{

	

	// 计算选择下一个城市的概率
	probabilities := make([]float64, len(unvisitedCities))
	totalProbability := 0.0

	for i, city := range unvisitedCities {
		if currentCity == city {
			continue;
		}else {
			pheromone := pheromoneMatrix[currentCity][city]
			heuristic := 1.0 / distanceMatrix[currentCity][city] // 距离的倒数作为启发式信息
			probabilities[i] = math.Pow(pheromone, Alpha) * math.Pow(heuristic, Beta)
			totalProbability += probabilities[i]
		}
	
	}

	// if len(unvisitedCities) == 1 && distanceMatrix[ant.CurrentPos][unvisitedCities[0]] == math.Inf(1){
	// 	ant.Tour=nil
	// 	ant.TotalCost=math.Inf(1)
	// 	ant.Visited[unvisitedCities[0]]=true
	// } else{

	
	if totalProbability > 0 {
		// 根据概率选择城市
		r := rand.Float64() * totalProbability
		accumulator := 0.0
		// for i, _ := range unvisitedCities {
		for i:=0; i<len(unvisitedCities); i++ {
			accumulator += probabilities[i]
			if accumulator >= r {
				selectedCity = unvisitedCities[i]
				break
			}
		}
	}
	

	//待更改
	return selectedCity

}

func UpdatePheromone(pheromoneMatrix [][]float64,distanceMatrix [][]float64, ants []*Ant) {
	// 在这里更新信息素浓度
	// 信息素挥发 - 对现有信息素进行衰减
	for i := range pheromoneMatrix {
		for j := range pheromoneMatrix[i] {
			pheromoneMatrix[i][j] *= (1.0 - Rho)
		}
	}

	// 更新信息素浓度 - 根据每只蚂蚁访问的路径
	for _, ant := range ants {
		if len(ant.Tour) > 1 || ant.Tour == nil {
			for i := 0; i < len(ant.Tour)-1; i++ {
				from := ant.Tour[i]
				to := ant.Tour[i+1]
				// 计算信息素增量
				pheromoneDelta := 5.0 / distanceMatrix[from][to] // 5乘以距离的倒数
				// 更新信息素浓度
				pheromoneMatrix[from][to] += pheromoneDelta
				pheromoneMatrix[to][from] = pheromoneMatrix[from][to] // 信息素浓度矩阵是对称的
			}
		}
	}
}

func findBestAnt(ants []*Ant) *Ant {
	// 找到最佳蚂蚁
	bestAnt := ants[0]
	for _, ant := range ants {
		if ant.TotalCost < bestAnt.TotalCost {
			bestAnt = ant
		}
	}
	return bestAnt
}