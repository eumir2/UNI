public class Punto2D {
    // OVERVIEW: classe che modella un punto IMMUTABILE in due dimensioni, `x` e
    // `y`, su numeri razionali.
    private final double x;
    private final double y;

// CONSTRUCTORS
    public Punto2D() {
    // MODIFIES: this
    // EFFECTS: istanzia un punto con coordinate nell'origine
        this.x = 0;
        this.y = 0;
    }

    public Punto2D(double x, double y) {
    // MODIFIES: this
    // EFFECTS: istanzia un punto con coordinate x e y passate come parametri
        this.x = x;
        this.y = y;
    }

// METHODS
    public double getX() {
        return x;
    }

    public double getY() {
        return y;
    }

    @Override
    public String toString() {
        return "Punto2D - (x:" + this.getX() + " , y: " + this.getY()+ ")"; 
    }

    @Override
    public boolean equals(Object obj) {
        if(obj instanceof Punto2D){
            Punto2D p = (Punto2D)obj;
            if((this.getX() == p.getX()) && (this.getY() == p.getY()))
                return true;
            return false;
        } 
        return false;
    }

    //repOK non necessaria

}
