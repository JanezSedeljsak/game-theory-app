export const GMEnum = { Multiplayer: 1, EasyAI: 2, AdvancedAI: 3 }
export const GEnum = { Connect4: 1, TicTacToe: 2 };

export function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms ?? 200));
}

export function GMEnumStr(gmenum) {
  switch (gmenum) {
    case GMEnum.AdvancedAI: return 'Advanced AI';
    case GMEnum.EasyAI: return 'Easy AI';
    case GMEnum.Multiplayer: return 'Multiplayer';
    default: throw new Error(`Invalid enum value - ${gmenum}!`)
  }
}

export async function boardAction(gameMode, mutateAI, mutateRand, multiplayer, ...args) {
  switch (gameMode) {
    case GMEnum.AdvancedAI: return await mutateAI(...args);
    case GMEnum.EasyAI: return await mutateRand(...args);
    case GMEnum.Multiplayer: return await multiplayer(...args);
    default: throw new Error(`Invalid gameMode enum - ${gameMode}!`);
  }
}

export function getSymbol(num) {
  switch (num) {
    case 1: return "O";
    case -1: return "X";
    case 0: return "";
    default: throw new Error(`Invalid num - ${num} (0,1,-1 are the only valid args)!`);
  }
}

function getPlayerLabels(game) {
  switch (game) {
    case GEnum.Connect4: return ['Green', 'Blue'];
    case GEnum.TicTacToe: return ['O', 'X'];
    default: throw new Error(`Invalid game enum - ${game}!`);
  }
}

export function getNextPlayer(xStart, playerMoveCount, gameMode) {
  if (gameMode != GMEnum.Multiplayer) return playerMoveCount % 2 == 0 ? 1 : -1;
  if (!xStart) return playerMoveCount % 2 == 0 ? 1 : -1;
  return playerMoveCount % 2 != 0 ? 1 : -1;
}

export function getNextRow(board, col) {
  let freeRow = 0
  while (board[freeRow][col] != 0) freeRow++;

  return freeRow;
}

export async function evalGameStatus(coreResponse, toasts, game) {
  const parsedResponse = JSON.parse(coreResponse);
  const response = [parsedResponse.board, new Set()];
  if (parsedResponse?.info && !parsedResponse.isdone) {
    toasts.clearAll();
    toasts.add({
      title: "Game status",
      description: parsedResponse.info,
      type: "info",
      duration: 10000,
    });
  }

  if (!parsedResponse.isdone) return response;
  const [p1, p2] = getPlayerLabels(game);
  const description = !parsedResponse.winner ? "Draw" : `${parsedResponse.winner == 1 ? p1 : p2} won!`;

  if (parsedResponse.winner != 0) {
    const n = game == GEnum.Connect4 ? 7 : 3;
    parsedResponse.coords.forEach(coord => response[1].add(coord.Row * n + coord.Col));
  }

  toasts.add({
    title: "Game finished",
    description: description,
    type: "info",
  });

  return response;
}

export async function histroyAction(prevMoveFunc) {
  const parsedResponse = JSON.parse(await prevMoveFunc());
  const response = [parsedResponse.board, new Set(), !parsedResponse.empty];
  return response
}
