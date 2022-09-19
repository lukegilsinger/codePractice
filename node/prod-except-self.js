var productExceptSelf = function(nums) {
  let r = []
  // revNums = nums.reverse()
  // console.log(revNums)
  nums.reduce((prev, acc, i) => {
      r[i] = prev
      return prev * acc
  }, 1)
  
  nums.reduceRight((prev, acc, i) => {
      r[i] *= prev
      return prev * acc
  }, 1)
  
  return r
};

const l = [1,2,3,4]
console.log(productExceptSelf(l))