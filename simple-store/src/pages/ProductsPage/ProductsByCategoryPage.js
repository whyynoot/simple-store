import ProductCard from "../../components/ProductCard/ProductCard";
import React, { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";
import CallRequestForm from '../../components/OrderRequest/OrderRequest';

function ProductsByCategoriesPage() {
    const [products, setProducts] = useState([]);
    let name = useParams();
  
    useEffect(() => {
      async function fetchProducts() { 
        const response = await fetch(`/api/category/${name.name}/products`);
        const data = await response.json();
        if (data.products.length === 0) {
          setProducts([])
          return ;
        }
        setProducts(data.products)
      }
      fetchProducts();
    }, []);
    
      return (
          <section className="py-5">
              <div className="container px-4 px-lg-5 mt-5">
                  <div className="row gx-4 gx-lg-5 row-cols-2 row-cols-md-3 row-cols-xl-4 justify-content-center">
                      {products.map(product => (
                          ProductCard(product)
                      ))}
                  </div>
              </div>
              <CallRequestForm />
          </section>
      );
  }

export default ProductsByCategoriesPage;