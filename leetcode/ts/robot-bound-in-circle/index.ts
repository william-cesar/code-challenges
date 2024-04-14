/*
 ends facing W -> returns home after 4 events
 ends facing S -> returns home after 2 events
 ends facing E - > returns home after 4 events
 ends facing N -> is circular if position ends in (0, 0). Otherwise it's gone  
*/

type Direction = 'N' | 'S' | 'E' | 'W';
type NextDirection = { [property: string]: Direction };
type NewPosition = { [property: string]: Function };

function isRobotBounded(instructions: string): boolean {
  const instList: Array<string> = instructions.split('');
  const position: Array<number> = [0, 0];
  let direction: Direction = 'N';

  const nextDirection: NextDirection = {
    LN: 'W',
    LW: 'S',
    LS: 'E',
    LE: 'N',
    RN: 'E',
    RE: 'S',
    RS: 'W',
    RW: 'N'
  };

  const newPosition: NewPosition = {
    N: () => (position[1] = position[1] + 1),
    S: () => (position[1] = position[1] - 1),
    E: () => (position[0] = position[0] + 1),
    W: () => (position[0] = position[0] - 1)
  };

  instList.forEach((instruction: string): void => {
    if (instruction === 'G') {
      newPosition[direction]();
    } else {
      direction = nextDirection[instruction + direction];
    }
  });

  return position.join('') === '00' || direction !== 'N';
}
