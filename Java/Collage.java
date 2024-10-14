import java.util.ArrayList;
import java.util.Scanner;


public class Collage {
//OVERVIEW: classe che modella un insieme di oggetti di tipo Figura
    private ArrayList<Figura> figure;

//CONSTRUCTOR
    public Collage(){
        this.figure = new ArrayList<Figura>();
    }

//METHODS
    public void ordinaPerimetro(){
        this.figure.sort(null);
    }

    public String toString(){
        String tmp = "Ordinati per perimetro: \n";
        for (Figura fTmp : this.figure) {
            tmp += fTmp.toString() + " - Perimetro: " + fTmp.perimetro()+ "\n";
        }
        return tmp;
    }

//MAIN
   public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        Collage c = new Collage();

        System.out.println("Inserisci delle figure (Termina con CTRL+D)");
        while(s.hasNext()){
            Figura f = null;
            String figura = s.next();
            if(figura.equals("Triangolo")){
                double lato1 = s.nextDouble();
                double lato2 = s.nextDouble();
                double lato3 = s.nextDouble();
                f = new Triangolo(lato1, lato2, lato3);

            }else if(figura.equals("Cerchio")){
                double raggio = s.nextDouble();
                f = new Cerchio(raggio);

            }else if(figura.equals("Rettangolo")){
                double base = s.nextDouble();
                double altezza = s.nextDouble();
                f = new Rettangolo(base, altezza);
            }else{
                double lato = s.nextDouble();
                f = new Quadrato(lato);
            }

            c.figure.add(f);
        }
        s.close();
        
        c.ordinaPerimetro();
        System.out.println(c.toString());
   } 
}
