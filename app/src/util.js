export function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms ?? 200));
}

export function transSymbol(symbol) {
    switch (symbol) {
      case 0:
        return "";
      case 1:
        return "O";
      case -1:
        return "X";
    }
  }

