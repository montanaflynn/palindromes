use std::io::{File, BufferedReader};
use std::iter::order;
use std::ascii::StrAsciiExt;

fn main() {
    let path = Path::new("/usr/share/dict/words");
    let mut file = BufferedReader::new(File::open(&path));
    for line in file.lines() {
        let line = line.unwrap();
        let lowercase = line.as_slice().to_ascii_lower();
        let word = lowercase.as_slice().trim_chars('\n');
        let forwards = word.chars();
        let backwards = word.chars().rev();
        if order::eq(forwards, backwards) {
            println!("{} is a palindrome", word);
        }
    }
}
