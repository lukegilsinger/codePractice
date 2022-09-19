var maxProfit = function(p) {
  let l = 0
  let r = 1
  let maxi = 0

  while(r < p.length) {
    if (p[l] < p[r]) {
      c = p[r] - p[l]
      maxi = Math.max(maxi, c)
    } else {
      l = r
    }
    r += 1
  } 
  return maxi
};

const prices = [7,3,5,1,6,4]
const r = maxProfit(prices)
console.log(r)