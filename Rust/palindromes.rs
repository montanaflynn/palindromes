use std::io::{File, BufferedReader};
use std::iter::order;
use std::ascii::AsciiExt;

fn main() {
    let path = Path::new("/usr/share/dict/words");
    let mut file = BufferedReader::new(File::open(&path));
    for line in file.lines() {
        let line = line.unwrap();
        let word = line.as_slice().trim_chars('\n');
        let lowercase = word.as_slice().to_ascii_lower();
        let forwards = lowercase.chars();
        let backwards = lowercase.chars().rev();
        if order::eq(forwards, backwards) {
            println!("{} is a palindrome", word);
        }
    }
}
