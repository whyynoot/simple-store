function ProductCard (props) {
    return (
    <div className="col mb-5">
        <div className="card h-100 border-success">
            <img className="card-img-top" src={"/image/" + props.image} alt="Image not found"/>
            <div className="card-body p-4">
                <div className="text-center">
                    <h5 className="fw-bolder">{props.name}</h5>
                    {props.price} рублей
                </div>
            </div>
            <div className="card-footer p-4 pt-0 border-top-0 bg-transparent">
                <div className="text-center"><a className="btn btn-outline-success mt-auto" href={`/product/${props.api_name}`}>Подробнее</a></div>
            </div>
        </div>
    </div>
    );
}

export default ProductCard;