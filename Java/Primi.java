import java.util.ArrayList;
import java.util.Iterator;

public class Primi {
    private ArrayList<Integer> numeri;

    // CONSTRUCTORS
    public Primi() {
        // MODIFIES: this
        // EFFECTS: modella un oggetto Primi con un'arrayList vuota
        this.numeri = new ArrayList<Integer>();
    }

    // METHODS
    public void generaNumeri(int n) {
        int i = 2;
        while (this.numeri.size() < n) {
            if(checkForPrime(i))
                this.numeri.add(i);
            i++;
        }
    }

    static boolean checkForPrime(int inputNumber){
        boolean isItPrime = true;

        if(inputNumber <= 1) {
            isItPrime = false;

            return isItPrime;
        }
        else{
            for (int i = 2; i<= inputNumber/2; i++) {
                if ((inputNumber % i) == 0){
                    isItPrime = false;
                    break;
                }
            }
            return isItPrime;
        }
    }

//MAIN
    public static void main(String[] args) {
        Primi p = new Primi();
        int n = Integer.parseInt(args[0]);
        p.generaNumeri(n);

        Iterator<Integer> i = p.numeri.iterator();
        while(i.hasNext())
            System.out.println(i.next() + " con iteratore"); 
    }
}