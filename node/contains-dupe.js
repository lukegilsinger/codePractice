var containsDuplicate = function(nums) {
    return new Set(nums).size !== nums.length
};

const l = [1,6,8,9,3,4,11,17,14,5]
console.log(containsDuplicate(l))