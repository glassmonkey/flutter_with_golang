//
//  api.swift
//  Runner
//
//  Created by 永野峻輔 on 2021/04/20.
//

import Foundation
import API


class Api: FlutterApi {
    func eval(_ input: FlutterRequest, error: AutoreleasingUnsafeMutablePointer<FlutterError?>) -> FlutterResponse? {
        let response = FlutterResponse()
        var result: Double = 0
        var error: NSError?
        ApiEvaluate(input.expression, &result, &error)
        if error != nil {
            response.message = "\(error)"
            return response
        }
        response.message = "\(result)"
        return response
    }
    
    func call(_ error: AutoreleasingUnsafeMutablePointer<FlutterError?>) -> FlutterResponse? {
        let response = FlutterResponse()
        var error: NSError?
        let text = ApiGoCall(&error)
        response.message = text
        return response
    }
}
