def quicksort array
  above = []
  below = []
  return array if array.size < 2
  median = array.delete_at(rand(array.size))
  array.each do |i|
    if i > median
      above << i
    else
      below << i
    end
  end

  return quicksort(below)+ [median] + quicksort(above)
end
