import random

arr = [random.randint(1, 999) for _ in range(100)]


def quicksort(array: list[int]):

    if len(array) <= 1:
        return array

    pivot_index = len(array) // 2
    pivot_value = array[pivot_index]

    left = list(filter(lambda x: x < pivot_value, array))
    right = list(filter(lambda x: x > pivot_value, array))

    return quicksort(left) + [pivot_value] + quicksort(right)


sorted_arr = quicksort(arr)
print(sorted_arr == sorted(sorted_arr))
