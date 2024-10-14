import java.util.ArrayList;
import java.util.Scanner;

public class Segmento2D implements Segmento{
//OVERVIEW: classe che modella un segmento in due dimensioni immutabile costituito da 2 Punto2D 
    private final Punto2D a;
    private final Punto2D b;

//CONSTRUCTORS
    public Segmento2D(Punto2D a, Punto2D b)throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: istanzia un Segmento2D con estremi 2 Punto2D presi come parametri
    //         se i due punti sono uguali o nulli lancia IllegalArgumentException
       
        if(a == null)
            throw new IllegalArgumentException("A non può essere null");
        if(b == null)
            throw new IllegalArgumentException("B non può essere null");  
        if(a.equals(b))  
            throw new IllegalArgumentException("I punti non possono essere uguali");

        this.a = a;
        this.b = b;
    }

//METHODS

    public Punto2D getA() {
        return a;
    }

    public Punto2D getB() {
        return b;
    }

    @Override
    public double length(){
    //EFFECTS: restituisce la lunghezza del punto
        return Math.sqrt(Math.pow(this.getA().getX() - this.getB().getX(),2) + Math.pow(this.getA().getY() - this.getB().getY(),2));
    }

    @Override
    public String toString() {
        return "Segmento2D -a: " + this.a.toString() + ", b: " + this.b.toString(); 
    }

//MAIN
    public static void main(String[] args) {
        double minLen = Double.parseDouble(args[0]);
        ArrayList<Segmento2D> segmenti = new ArrayList<Segmento2D>();

        Scanner s = new Scanner(System.in);

        System.out.println("Inserisci i segmenti nel formato ax ay bx by (temina con CTRL+D)");
        while(s.hasNext()){
            double ax = Double.parseDouble(s.next());
            double ay = Double.parseDouble(s.next());
            double bx = Double.parseDouble(s.next());
            double by = Double.parseDouble(s.next());

            Punto2D a = new Punto2D(ax,ay);
            Punto2D b = new Punto2D(bx,by);

            try {
                segmenti.add(new Segmento2D(a, b));
            } catch (IllegalArgumentException e) {
                System.out.println(e.getMessage());
            }
        }
        s.close();

        System.out.println("Segmenti di lunghezza superiore a " + minLen);
        for (Segmento2D segmento2d : segmenti) {
            if(segmento2d.length() > minLen)
                System.out.println(segmento2d.toString()+ "lunghezza: " + segmento2d.length());
        }
    }

}
