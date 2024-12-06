WITH ProductPriceDetails AS (
    SELECT 
        p.product_id, 
        p.price, 
        COALESCE(us.units, 0) AS units,
        p.start_date,
        p.end_date
    FROM Prices p
    LEFT JOIN UnitsSold us ON 
        p.product_id = us.product_id AND 
        us.purchase_date BETWEEN p.start_date AND p.end_date
)

SELECT 
    product_id, 
    COALESCE(
        ROUND(
            SUM(price * units) / 
            NULLIF(SUM(units), 0), 
        2), 
    0) AS average_price
FROM ProductPriceDetails
GROUP BY product_id;