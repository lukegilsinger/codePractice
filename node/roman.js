/**
 * @param {string} s
 * @return {number}
 */

 const characters = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
}
var romanToInt = function(s) {
  const chars = s.split("")
  const vals = chars.map((c) => characters[c])
  console.log(vals)
  let t = 0;
  vals.forEach((v, i) => {
    console.log(v)
    // console.log(vals[i+1])
    if (v >= vals[i+1] || vals[i+1] === undefined) {
      // console.log('yes')
      t = t + v;
    } else {
      // console.log("no")
      t = t - v;
    }
    console.log(t)
  });
  console.log(t)
};

// romanToInt("LVIII")
romanToInt("MCMXCIV")