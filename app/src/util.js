export function getSymbol(symbol) {
  switch (symbol) {
    case 0:
      return "";
    case 1:
      return "O";
    case -1:
      return "X";
  }
}

export const GMEnum = {
  Multiplayer: 1,
  EasyAI: 2,
  AdvancedAI: 3
}

export function GMEnumStr(gmenum) {
  switch (gmenum) {
    case GMEnum.AdvancedAI:
      return 'Advanced AI';
    case GMEnum.EasyAI:
      return 'Easy AI';
    case GMEnum.Multiplayer:
      return 'Multiplayer';
    default:
      throw new Error(`Invalid enum value - ${gmenum}!`)
  }
}

export function getNextPlayer(xStart, playerMoveCount, gameMode) {
  if (gameMode != GMEnum.Multiplayer) {
    return playerMoveCount % 2 == 0 ? 1 : -1;
  }

  if (!xStart) {
    return playerMoveCount % 2 == 0 ? 1 : -1;
  }

  return playerMoveCount % 2 != 0 ? 1 : -1;
}

export function getNextRow(board, col) {
  let freeRow = 0
  while (board[freeRow][col] != 0) {
    freeRow++;
  }

  return freeRow;
}
