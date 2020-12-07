package main

type FallPath struct {
	Right int
	Down  int
}

type Slope interface {
	Right(count int)
	Down(count int)
	AtBottom() bool
	HaveHitTree() bool
}

func NewPathMapper(slope Slope) *PathMapper {
	return &PathMapper{
		slope: slope,
	}
}

type PathMapper struct {
	fallPath   FallPath
	slope      Slope
	treesHit   int
	calculated bool
}

func (p *PathMapper) CalculateTreesHit(fallPath FallPath) int {
	if p.calculated {
		return p.treesHit
	}

	for {
		p.slope.Right(fallPath.Right)
		p.slope.Down(fallPath.Down)

		if p.slope.AtBottom() {
			break
		}

		if p.slope.HaveHitTree() {
			p.treesHit++
		}
	}

	p.calculated = true
	return p.treesHit
}

func (p *PathMapper) Reset() {
	p.treesHit = 0
	p.calculated = false
}
