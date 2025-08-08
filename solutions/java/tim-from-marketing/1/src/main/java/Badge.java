class Badge {
    public String print(Integer id, String name, String department) {
        department = department != null ? department : "owner";
        if (id == null) {
            return String.format("%s - %S", name, department);
        } else {
            return String.format("[%d] - %s - %S", id, name, department);
        }
    }
}
