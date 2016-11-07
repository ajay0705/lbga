package main

import (
	"math/rand"
	"fmt"
)

type task struct {
	task_no float64
	duration int32
	status int
}

type chromosome struct {
	tasks []task
	machine_no int
	fitness float64
}

type constants struct {
	crossover_ratio float64
	mutation_ratio float64
	elitism_ratio float64
	popsize int
	chrom_length int
	max_generations int
}

func main() {
	var constant constants
	constant.crossover_ratio=0.6
	constant.mutation_ratio=0.05
	constant.elitism_ratio=0.5
	/* constant.popsize=
	constant.chrom_length=
	constant.max_generations=
	*/

}

func perm_of_tasks() {
	/* 1. Decide on which tasks will be chosen for allocation
	   2. Given them status 0 for unallocated, 1 for allocated and 2 for completed
	   If tasks are completed, then dequeue them*/
}

func gen_population (totalperms []chromosome, k int) []chromosome {
	var population []chromosome
	var i int
	var index int
	for i=0;i<k;i++ {
		index=rand.Intn(k)
		append(population, totalperms[index])
	}
	return population
}

func calculate_fitness(population []chromosome) {
	var i int
	for i=0; i<constants.popsize; i++ {
		population[i].fitness=0 /* formula  WRITE THE FORMULA FOR FITNESS FUNCTION */
	}
}

func roulettewheel(popsize int, sumfitness float64, population []chromosome) int {
	var randno, partialsum float64
	var j int32
	partialsum=0.0
	randno=rand.Float32()*sumfitness
	for j=0; j < popsize; j++ {
		partialsum:=partialsum+population[j].fitness
		if partialsum >= randno {
			break
		}
	}
	return j
}

func roulettewheelsa(popsize int, population []chromosome) []int {
	maxfitness := maxno(population, constants.popsize)
	var counter []int
	n_select := 4000 /* IS THIS THE CORRECT NUMBER AND WHAT IS n_select? */
	var i int32
	var notaccepted bool
	for i = 0; i < n_select; i++ {
		notaccepted = true
		for (notaccepted == true) {
			index := rand.Intn(popsize)
			if rand.Float64() < (population[index].fitness / maxfitness) {
				notaccepted = false
			}
			counter[i]++
		}

	}
	/* Printing all the counters coz i don't know what to return from this function */

	for i=0;i<popsize;i++ {
		fmt.Println("counter[" + i + "]:" + counter[i])
	}
	return counter
}

func tournamentselect(popsize int, population []chromosome, k int) int {
	/* k should not be equal to popsize, but rather a number less than that. Please find out how to fix k */
	best:=0
	var i int
	var index int
	for i=0; i<k; i++ {
		index=rand.Intn(popsize)
		if (best == 0) || population[index].fitness > population[best].fitness {
			best=index
		}

	}
	return best
}

func crossover(parent1 chromosome, parent2 chromosome) []chromosome {
	var pivot int
	pivot = rand.Float32() * len(parent1.tasks)
	var first, second []task
	for i:=0; i< len(parent1.tasks); i++ {
		if i < pivot {
			first = append(first, parent1.tasks[i])
			second = append(second, parent2.tasks[i])
		} else {
			first = append(first, parent1.tasks[i])
			second = append(second, parent2.tasks[i])
		}
	}
	return []chromosome{{parent1: first}, {parent2: second}}
}

func mutation(child chromosome) chromosome {
	/* Out of all the types of mutation like bit flip, random resetting, swap mutation etc,
	we are using swap mutation
	 */
	l := len(child.tasks)
	pivot1 := rand.Intn(l)
	pivot2 := rand.Intn(l)
	temp := child.tasks[pivot1]
	child.tasks[pivot1] = child.tasks[pivot2]
	child.tasks[pivot2] = temp

	return child
}

func maxno(population []chromosome, popsize int) int {
	var biggest=population[0].fitness
	var i int
	var curr int
	for i=0;i<popsize;i++ {
		curr=population[i].fitness
		if curr > biggest {
			biggest = curr
		}
	}
	return curr
}

