import java.io.File;
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.IOException;
 
class Palindromes
{
  public static void main(String args[]) throws IOException 
  {
    File file = new File("/usr/share/dict/words");
    FileReader fr = new FileReader(file);
    BufferedReader br = new BufferedReader(fr);
    String line;
    while((line = br.readLine()) != null){
      String word = line;
      String forwards = word.toLowerCase();
      String backwards = new StringBuffer(forwards).reverse().toString();
      if (forwards.equals(backwards))
        System.out.println(String.format("%s is a palindrome", word));
    }
    br.close();
    fr.close();
  }
}
