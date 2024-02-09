package day5

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const SPVERSION1 = 1
const SPVERSION2 = 2

type Almanac struct {
	Seeds                    []int
	SeedLocations            []int
	SeedToSoilMap            TranslationMap
	SoilToFertilizerMap      TranslationMap
	FertilizerToWaterMap     TranslationMap
	WaterToLightMap          TranslationMap
	LightToTemperatureMap    TranslationMap
	TemperatureToHumidityMap TranslationMap
	HumidityToLocationMap    TranslationMap
}

type TranslationMap struct {
	Ranges []int
	Offset []int
}

func (t *TranslationMap) AddValue(src, rng, dest int) {
	if len(t.Ranges) == 0 {
		t.Ranges = []int{0}
		t.Offset = []int{0}
	}

	i := sort.SearchInts(t.Ranges, src)
	if i < len(t.Ranges) && t.Ranges[i] == src {
		t.Offset[i] = dest - src
	} else {
		t.Ranges = append(t.Ranges[:i], append([]int{src}, t.Ranges[i:]...)...)
		t.Offset = append(t.Offset[:i], append([]int{dest - src}, t.Offset[i:]...)...)
	}

	if i+1 < len(t.Ranges) && t.Ranges[i+1] == src+rng {
		return
	}

	t.Ranges = append(t.Ranges[:i+1], append([]int{src + rng}, t.Ranges[i+1:]...)...)
	t.Offset = append(t.Offset[:i+1], append([]int{0}, t.Offset[i+1:]...)...)
}

func (t *TranslationMap) GetOffset(i int) int {
	oi := sort.SearchInts(t.Ranges, i)
	if oi < len(t.Ranges) && t.Ranges[oi] == i {
		return t.Offset[oi]
	}
	return t.Offset[oi-1]
}

func parseMap(lines []string) TranslationMap {
	m := TranslationMap{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dest, _ := strconv.Atoi(parts[0])
		src, _ := strconv.Atoi(parts[1])
		rng, _ := strconv.Atoi(parts[2])

		m.AddValue(src, rng, dest)
	}

	return m
}

func (a *Almanac) parseMaps(maps [7][]string) {
	var wg sync.WaitGroup

	wg.Add(7)
	go func() {
		defer wg.Done()
		a.SeedToSoilMap = parseMap(maps[0])
	}()
	go func() {
		defer wg.Done()
		a.SoilToFertilizerMap = parseMap(maps[1])
	}()
	go func() {
		defer wg.Done()
		a.FertilizerToWaterMap = parseMap(maps[2])
	}()
	go func() {
		defer wg.Done()
		a.WaterToLightMap = parseMap(maps[3])
	}()
	go func() {
		defer wg.Done()
		a.LightToTemperatureMap = parseMap(maps[4])
	}()
	go func() {
		defer wg.Done()
		a.TemperatureToHumidityMap = parseMap(maps[5])
	}()
	go func() {
		defer wg.Done()
		a.HumidityToLocationMap = parseMap(maps[6])
	}()

	wg.Wait()
}

func parseSeedsV1(seeds []string) []int {
	intSeeds := []int{}
	for _, seed := range seeds {
		intSeed, _ := strconv.Atoi(seed)
		intSeeds = append(intSeeds, intSeed)
	}

	return intSeeds
}

func parseSeedsV2(seeds []string) []int {
	intSeeds := []int{}
	for i := 0; i < len(seeds); i += 2 {
		startSeed, _ := strconv.Atoi(seeds[i])
		seedRange, _ := strconv.Atoi(seeds[i+1])
		for seedNum := startSeed; seedNum < seedRange+startSeed; seedNum += 1 {
			intSeeds = append(intSeeds, seedNum)
		}
	}

	return intSeeds
}

func parseAlmanac(lines []string, spVersion int) (Almanac, error) {
	a := Almanac{}

	switch spVersion {
	case 1:
		a.Seeds = parseSeedsV1(strings.Split(strings.Split(lines[0], ": ")[1], " "))
	case 2:
		a.Seeds = parseSeedsV2(strings.Split(strings.Split(lines[0], ": ")[1], " "))
	default:
		return a, fmt.Errorf("unknown spVersion: %v", spVersion)
	}

	breaks := []int{}
	for i, line := range lines {
		if line == "" {
			breaks = append(breaks, i)
		}
	}
	breaks = append(breaks, len(lines))

	maps := [7][]string{}
	for i := 0; i < len(breaks)-1; i++ {
		start := breaks[i] + 2
		stop := breaks[i+1]
		maps[i] = lines[start:stop]
	}

	a.parseMaps(maps)

	return a, nil
}

func (a *Almanac) findSeedLocation(i int) {
	seed := a.Seeds[i]
	seed += a.SeedToSoilMap.GetOffset(seed)
	seed += a.SoilToFertilizerMap.GetOffset(seed)
	seed += a.FertilizerToWaterMap.GetOffset(seed)
	seed += a.WaterToLightMap.GetOffset(seed)
	seed += a.LightToTemperatureMap.GetOffset(seed)
	seed += a.TemperatureToHumidityMap.GetOffset(seed)
	seed += a.HumidityToLocationMap.GetOffset(seed)
	a.Seeds[i] = seed
}

func (a *Almanac) findSeedLocations() {
	for i := range len(a.Seeds) {
		go a.findSeedLocation(i)
	}
}

func (a *Almanac) GetLowestSeedLocation() int {
	a.findSeedLocations()

	lowest := int(math.Inf(1))
	for _, seed := range a.Seeds {
		if seed < lowest {
			lowest = seed
		}
	}

	return lowest
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day5/input/p1.txt")
	if err != nil {
		return nil, err
	}

	a, err := parseAlmanac(lines, SPVERSION1)
	if err != nil {
		return nil, err
	}

	return a.GetLowestSeedLocation(), nil
}

func part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day5/input/p1.txt")
	if err != nil {
		return nil, err
	}

	a, err := parseAlmanac(lines, SPVERSION2)
	if err != nil {
		return nil, err
	}

	return a.GetLowestSeedLocation(), nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
