import java.net.SocketTimeoutException;
import java.util.ArrayList;
import java.util.Scanner;

public class Sheep implements Cloneable{
//OVERVIEW: definita da un nome e da un attributo `dna`, contenente una serie di stringhe che definiscono i codici dna dei diversi cromosomi della pecora.
    public String nome;
    public ArrayList<String> dna;

    public static void main(String[] args) {
        Sheep first = new Sheep(args[0]);
        System.out.println("Inserisci i cromosomi, uno per riga (terminare con CTRL+D)");
        
        Scanner s = new Scanner(System.in);
        while(s.hasNext()){
            String tmp = s.next();
            first.dna.add(tmp);   
        }
        //s.close();

        //System.out.println(first);

        Sheep clone = (Sheep)first.clone();
        System.out.println("Ho creato un clone di "+ first.nome);

        if(first.equals(clone))
            System.out.println("Le due pecore sono uguali");
        else
            System.out.println("Le due pecore sono diverse");


        //s.useDelimiter(System.lineSeparator());
        System.out.println("Inserisci il cromosoma da modificare ed il nuovo codice:");
       
        /*s.nextLine();
        int indice = Integer.parseInt(s.nextLine());
        String codice = s.next();*/

        int indice = 3;
        String codice = "CATAGATAGATAGAG";

        clone.sostituisciDna(indice, codice);

        if(first.equals(clone))
            System.out.println("Le due pecore sono uguali");
        else
            System.out.println("Le due pecore sono diverse");


    }
    
//CONSTRUCTORS
    public Sheep(String nome) {
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo Sheep
        this.nome = nome;
        this.dna = new ArrayList<String>();
    }

//METHODS
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
    public ArrayList<String> getDna() {
        return dna;
    }
    public void setDna(ArrayList<String> dna) {
        this.dna = dna;
    }

    public String getStringaDna(int index){
        return dna.get(index);
    }

    public void sostituisciDna(int index, String newDna){
        dna.set(index,newDna);
    }

    @Override
    protected Object clone(){
        Sheep clone = null;
        
        try{
            clone = (Sheep)super.clone();
            clone.dna = new ArrayList(this.dna);
        }catch(CloneNotSupportedException e){
            System.out.println("Errore nella clonazione");
        }

        return clone;
    }

    @Override
    public boolean equals(Object obj) {
        if(obj == null)
            return false;
        if(!(obj instanceof Sheep))
            return false;
        
        Sheep s = (Sheep)obj;
        if(!(this.nome.equals(s.nome)))
            return false;

        if(!(this.dna.equals(s.dna)))
            return false;
            
        return true;
    }
  

}
