// learn better for loop in TS

var solution = function (isBadVersion: any) {
  return function (n: number): number {
    for (let i = 0; i < n; i++) {
      if (isBadVersion(i)) {
        return i;
      }
    }
    return n;
  };
};

var solution = function (isBadVersion: any) {
  return function (n: number): number {
    let i = 0;
    while (i < n) {
      if (isBadVersion(i)) {
        return i;
      }
      i++;
    }
    return n;
  };
};
