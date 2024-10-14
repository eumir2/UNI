import java.time.LocalDate;

public class Dramma extends Film{


//CONSTRUCTOR
    public Dramma(String titolo, String tipo, LocalDate dN) {
        super(titolo, tipo, dN);
    }

    @Override
    public double calcolaPenaleRitardo() {
        return super.calcolaPenaleRitardo();
    }
        
}
