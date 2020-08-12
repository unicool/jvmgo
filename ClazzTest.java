

public class ClazzTest {
    public static final boolean FLAG = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public static final int INT = 123456789;
    public static final long LONG = 12345678901L;
    public static final float PI = 3.14f;
    public static final double E = 2.71828;

    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) throws RuntimeException {
        int x = 32760; // ldc

        ClazzTest.staticVar = x + 1; // putstatic
        x = ClazzTest.staticVar + 1; // getstatic

        ClazzTest obj = new ClazzTest(); // new
        obj.instanceVar = x + 1; // putfield
        x = obj.instanceVar + 1; // getfield
        Object o2 = obj;
        if (o2 instanceof ClazzTest) { // instanceof
            obj = (ClazzTest) o2; // checkcast
            System.out.println(obj.instanceVar);
        }
    }
}
