/**
 * Busca dentro del archivo json el registro con el id proporcionado
 * @param {Number} id ID de la ruta que se requiere buscar
 * @returns Objeto con la información necesaria
 */
function getRouteById(id) {
    let data = require("./routes.json")
    return data.routes.filter(i => i.id == id)[0]
}

/**
 * Toma la información del id encontrado y lo regresa dentro de un objeto con la data recibida
 * @param {Number} id Id de la ruta a buscar
 * @param {Object} data Información que se desea agregar
 * @returns Objeto con la información completa de la ruta
 */
function getDataOnFormat(id, data) {
    const stored = getRouteById(id)
    return {
        ...stored,
        ...data
    }
}

module.exports = {
    getDataOnFormat,
    getRouteById
}