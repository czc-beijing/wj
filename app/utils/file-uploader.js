import { wxToPromise } from "./wx";
import APIConfig from "../config/api";
import Http from "./http";

/**
 * @Created by WebStorm
 * @Author: 沁塵
 * @Desc:
 */

class FileUploader extends Http {
    static async upload(filePath, key = 'file') {
        let res
        try {
            res = await wxToPromise('uploadFile', {
                url: APIConfig.baseUrl + '/v1/file',
                filePath: filePath,
                name: key,
            })
        } catch (e) {
            FileUploader._showError(-1)
            throw new Error(e.errMsg)
        }

        const serverData = JSON.parse(res.data)

        if (res.statusCode !== 201) {
            FileUploader._showError(serverData.error_code, serverData.message)
            throw new Error(serverData.message)
        }

        return serverData.data
    }
    static async onlineUpload(filePath, key = 'file') {
        let res
        try {
            res = await wxToPromise('uploadFile', {
                url: 'https://qinchenju.com/homemaking/v1/file',
                filePath: filePath,
                name: key,
            })
        } catch (e) {
            FileUploader._showError(-1)
            throw new Error(e.errMsg)
        }

        const serverData = JSON.parse(res.data)

        if (res.statusCode !== 201) {
            FileUploader._showError(serverData.error_code, serverData.message)
            throw new Error(serverData.message)
        }

        return serverData.data
    }
}

export default FileUploader