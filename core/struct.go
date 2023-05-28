package core

const (
	END_FLAG       uint8 = 0b10000000
	RANGE_FLAG     uint8 = 0b01000000
	SUB_START_FLAG uint8 = 0b00100000
	SUB_END_FLAG   uint8 = 0b00010000
)

type State struct {
	Next  *State
	State uint8
	Range [][2]byte
	Set   []byte
}

func (s *State) IsEnd() bool {
	return s.State&END_FLAG == END_FLAG
}

func (s *State) IsRange() bool {
	return s.State&RANGE_FLAG == RANGE_FLAG
}

func (s *State) IsSubStart() bool {
	return s.State&SUB_START_FLAG == SUB_START_FLAG
}

func (s *State) IsSubEnd() bool {
	return s.State&SUB_END_FLAG == SUB_END_FLAG
}

type Regexp struct {
	Entry *State
	subs  [][]byte
}

func (re *Regexp) reset() {

}

func (re *Regexp) doMatch(data []byte) {
	cState := re.Entry
	buffer := []byte{}
L1:
	for i, b := range data {
		if cState.IsRange() {
			for _, r := range cState.Range {
				if b > r[0] && b < r[1] {
					break L1
				}
			}
		}
	}
}
