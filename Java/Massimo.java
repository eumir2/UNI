import java.util.ArrayList;
import java.util.Comparator;
import java.util.Scanner;

public class Massimo<E>{
//OVERVEW: classe che modella un generico insieme di elementi
    private ArrayList<E> elementi;
    //private E max;

//CONSTRUCTOR
    public Massimo(){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto Massimo vuoto
        this.elementi = new ArrayList<E>();
    }

//METHODS
    public void insert(E elem) throws IllegalArgumentException{
        if(elem == null)
            throw new IllegalArgumentException("Elemento nullo");
        else
            this.elementi.add(elem);
    }

    @Override
    public String toString() {
        String tmp = "Elementi:\n";
        for (E e : elementi) {
            tmp += e + "\n";
        }
        return tmp;
    }

//MAIN
    public static void main(String[] args) {
        Massimo m = new Massimo();

        Scanner s = new Scanner(System.in);
        System.out.println("Inserisci cose:");
        while(s.hasNext()){
            m.elementi.add(s.next());
        }

        System.out.println(m.toString());

        m.elementi.sort(null);
        System.out.println("Comparator null: " + m.elementi.get(m.elementi.size() - 1));

        Comparator c = new Comparator<String>() {

            @Override
            public int compare(String o1, String o2) {
                if(o1.length() < o2.length())
                    return -1;
                else if(o1.length() > o2.length())
                    return 1;
                else
                    return 0;
            }   
        };
        m.elementi.sort(c);

        System.out.println("Comparator definito nella classe: " + m.elementi.get(m.elementi.size() - 1));
        
    }

    


}