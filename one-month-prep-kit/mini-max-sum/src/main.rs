fn mini_max_sum(arr: Vec<i64>) {
    let mut min: i64 = 0;
    let mut max: i64 = 0;
    let mut items = arr;

    items.sort_unstable();

    for (index, num) in items.iter().enumerate() {
        if index != 0 {
            max = max + num;
        }
        if index != items.len() - 1 {
            min = min + num
        }
    }

    println!("{} {}", min, max);
}

fn main() {
    let arr = vec![7, 69, 2, 221, 8974];

    mini_max_sum(arr)
}
