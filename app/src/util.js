export function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms ?? 200));
}

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

