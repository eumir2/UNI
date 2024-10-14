public class Punto3D{
//OVERVIEW: classe che modella un punto immutabile in 3 dimensioni che estende Punto2D con aggiunta di coordinata z
    private final Punto2D p; //composizione
    private final double z;


//CONSTRUCTORS
    public Punto3D(double x, double y, double z){
        this.p = new Punto2D(x,y);
        this.z = z;
    }


    public double getZ(){
        return this.z;
    }

    //delega dei metodi al Punto2D
    public double getX(){
        return p.getX();
    }

    
    public double getY(){
        return p.getY();
    }

    @Override
    public boolean equals(Object obj) {
    if(obj instanceof Punto3D){
        Punto3D p = (Punto3D)obj;
        if((this.getX() == p.getX()) && (this.getY() == p.getY()) && (this.getZ() == p.getZ()))
            return true;
        }
        return false;
    }

    @Override
    public String toString() {
        return "Punto3D - (x:" + p.getX() + " , y: " + p.getY()+ ", z: " + this.getZ() + ")"; 
    }
}
