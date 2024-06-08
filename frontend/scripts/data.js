
const URL_BASE="http://localhost:3000/";

const endpoints ={
    users:`${URL_BASE}users`,
    getAnUserLogin:(usu) => `${URL_BASE}users?username=${usu}`,
    automobiles: `${URL_BASE}automobiles`,
    reservas: `${URL_BASE}reservas`
}

export default endpoints;