import java.util.ArrayList;
import java.util.Comparator;
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

    public void ordinaPerArea(){
        Comparator<Figura> c = new Comparator<Figura>() {

            @Override
            public int compare(Figura o1, Figura o2) {
                if(o1.area() > o2.area())
                    return -1;
                else if(o1.area() < o2.area())
                    return 1;
                else
                    return 0;
            }
            
        };
        this.figure.sort(c);
    }

    public double areaTotale(){
        double  tmp = 0.0;
        for (Figura fTmp : this.figure) {
            tmp += fTmp.area();
        }
        return tmp;
    }

    public String toString(){
        String tmp = "";
        for (Figura fTmp : this.figure) {
            tmp += fTmp.toString() + " Area: " + fTmp.area() + " - Perimetro: " + fTmp.perimetro()+ "\n";
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
        
        System.out.println("Area totale: "+ (c.areaTotale()));

        c.ordinaPerArea();
        System.out.println("Ordinati per area:\n" + c.toString());

        System.out.println("\n---------------\n");

        c.ordinaPerimetro();
        System.out.println("Ordinati per perimetro: \n " + c.toString());
   } 
}
