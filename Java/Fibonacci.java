import java.util.ArrayList;
import java.util.Iterator;

public class Fibonacci{
    public ArrayList<Integer> sequenza;

//CONSTRUCTORS
    public Fibonacci(){
        this.sequenza = new ArrayList<Integer>();
    }

//METHODS
    public void generaSequenza(int n){
        this.sequenza.add(1);
        this.sequenza.add(1);

        int num1 = 1;
        int num2 = 1;
        int somma = 0;
        while(n > 2){
            somma = num1+ num2;
            this.sequenza.add(somma);
            num1 = num2;
            num2 = somma;
            n--;
        }
    }

    /*   @Override
    public String toString() {
        String tmp= "";
        for (Integer integer : sequenza) {
            tmp += integer + "\n";
        }
        return tmp;
    }
    */
   
//MAIN
    public static void main(String[] args) {
        int n = Integer.parseInt(args[0]);

        Fibonacci f = new Fibonacci();
        f.generaSequenza(n);

        Iterator<Integer> i = f.sequenza.iterator();

        while(i.hasNext()){
            System.out.println(i.next() + " con iteratore");
        }

        //System.out.print(f.toString());

    }



}