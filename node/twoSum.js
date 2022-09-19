/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
  const hm ={};
  for (const [i, n] of nums.entries()) {
    const diff = target - n
    console.log(hm[diff] !== undefined)
    console.log(hm[diff])
    if(hm[diff] !== undefined) {
      console.log(diff)
      console.log([hm[diff], i])
      return [hm[diff], i]
    }
    // hm.set(n,i)
    hm[n] = i
  }
};

const nums = [2,7,11,15]
const target = 9

const ans = twoSum(nums, target)
console.log(ans)
