import { useEffect, useState } from "react";
import axios from "axios";

const Overview = () => {

    const [refreshTime, setRefreshTime] = useState("10");
    const [isRefreshing, setIsRefreshing] = useState(false);

    useEffect(() => {
        setIsRefreshing(true);

        axios
            .get("http://localhost:8080/api")
            .then((response) => {
                console.log(response);
            })
            .catch((error) => {
                console.log(error);
            })

        setIsRefreshing(false);
    }, refreshTime)

    return (
        <div>
            ffffdsa
        </div>
    );
}

export default Overview;