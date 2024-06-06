
const URL_BASE="http://localhost:3000/";

const endpoints ={
    users:`${URL_BASE}users`,
    getAnUserLogin:(usu) => `${URL_BASE}users?Usu=${usu}`,
    automobiles: `${URL_BASE}automobiles`
}

export default endpoints;