package nagano.shunsuke.flutter_with_golang


import api.Api
import nagano.shunsuke.plugins.Pigeon

class Api() : Pigeon.Api {
    override fun call(): Pigeon.Response {
        val response = Pigeon.Response()
        response.message = Api.goCall()
        return response
    }

    override fun eval(arg: Pigeon.Request): Pigeon.Response {
        val response = Pigeon.Response()
        return try {
            val result = Api.evaluate(arg.expression)
            response.message = result.toString()
            response
        } catch (e: Exception ) {
            response.message = e.toString()
            response
        }
    }
}