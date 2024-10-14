public class Segmento3D implements Segmento {
    private final Punto3D a;
    private final Punto3D b;
    private final Punto3D c;

//CONSTRUCTORS
    public Segmento3D(Punto3D a, Punto3D b, Punto3D c)throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: istanzia un Segmento2D con estremi 2 Punto2D presi come parametri
    //         se i due punti sono uguali o nulli lancia IllegalArgumentException
       
        if(a == null)
            throw new IllegalArgumentException("A non può essere null");
        if(b == null)
            throw new IllegalArgumentException("B non può essere null");  
        if(c == null)
            throw new IllegalArgumentException("C non può essere null");
        if(a.equals(b) && b.equals(c))  
            throw new IllegalArgumentException("I punti non possono essere uguali");

        this.a = a;
        this.b = b;
        this.c = c;
    }

//METHODS

    public Punto3D getA() {
        return a;
    }

    public Punto3D getB() {
        return b;
    }


    public Punto3D getC() {
        return c;
    }

    @Override
    public double length(){
    //EFFECTS: restituisce la lunghezza del punto
    double dx = Math.pow(this.getA().getX() - this.getB().getX(),2);
    double dy = Math.pow(this.getA().getY() - this.getB().getX(),2);
    double dz = Math.pow(this.getA().getZ() - this.getB().getZ(),2);
        return Math.sqrt(dx + dy + dz);
    }

    @Override
    public String toString() {
        return "Segmento3D -a: " + this.a.toString() + ", b: " + this.b.toString()  + ", c: " + this.c.toString(); 
    }

//MAIN
    public static void main(String[] args) {
    }
}
