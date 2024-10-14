
import java.util.ArrayList;
import java.util.Iterator;



public class Alfabeto {

    private ArrayList<Character> lettere;

//CONSTRUCTORS
    public Alfabeto(){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto alfabeto contenente tutte le lettere minuscole
    this.lettere = new ArrayList<Character>();
        for(int i = 97; i <= 122; i++){
            this.lettere.add((char)i);
        }
    }

    public Alfabeto(char inizio, char fine)throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto alfabeto contenente le lettere comprese tra quelle passate come parametro
        if(inizio < 97)
            throw new IllegalArgumentException("Valore limite iniziale non valido");
        if(fine > 122)
            throw new IllegalArgumentException("Valore limite finale non valido");
        if(inizio > fine)
        throw new IllegalArgumentException("La prima lettera è successiva alla seconda");

        this.lettere = new ArrayList<Character>();
        for(int i = 97; i <= 122; i++){
            if(i >= inizio && i <= fine){
                this.lettere.add((char)i);
            }
        }
    }

//METHODS
    public String toString(){
        String tmp = "";
        for (char c : this.lettere) {
            tmp += c + "\n";
        }
        return tmp;
    }




//MAIN
    public static void main(String[] args) {

        Alfabeto a;
        if(args.length > 0){
            char inizio = args[0].charAt(0);
            char fine = args[1].charAt(0);
            a = new Alfabeto(inizio,fine);
        }else{
            a = new Alfabeto();
        }
        

        Iterator<Character> i = a.lettere.iterator();
        while(i.hasNext()){
            System.out.print(i.next() + " con iteratore \n");
        }


    }
}
