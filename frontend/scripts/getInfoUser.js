
import endpoints from "./data.js";

// Hay que ingresar usu
export const UserLogin = async (usu) => {
  try {
    const { data } = await axios.get(endpoints.getAnUserLogin(usu));
    return data;
  } catch (error) {
    console.error(error);
    return [];
  }
};

export const getAllUsers = async () => {
    try {
      const { data } = await axios.get(endpoints.users);
      return data;
    } catch (error) {
      console.error(error);
      return [];
    }
};

export const getAutomobiles = async () => {
  try {
    const { data } = await axios.get(endpoints.automobiles);
    return data;
  } catch (error) {
    console.error(error);
    return [];
  }
};
  