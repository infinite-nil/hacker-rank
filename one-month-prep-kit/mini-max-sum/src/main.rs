fn mini_max_sum(arr: Vec<i32>) {
    let mut min: i32 = 0;
    let mut max: i32 = 0;

    for (index, num) in arr.iter().enumerate() {
        if index != 0 {
            max = max + num
        }
        
        if index != arr.len() - 1 {
            min = min + num
        }
    }

    println!("{} {}", min, max);
}

fn main() {
    let arr = vec![1, 3, 5, 7, 9];

    mini_max_sum(arr)
}