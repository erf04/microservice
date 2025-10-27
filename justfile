set shell := ["cmd.exe", "/C"]

order_build_dir := "F:/projects/kubernetes-test/order"
product_build_dir := "F:/projects/kubernetes-test/product"
run-order:
    cd {{order_build_dir}} && uvicorn main:app --host 0.0.0.0 --port 8000


run-product:
    cd {{product_build_dir}} && air