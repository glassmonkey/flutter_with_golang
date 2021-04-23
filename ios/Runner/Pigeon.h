// Autogenerated from Pigeon (v0.2.1), do not edit directly.
// See also: https://pub.dev/packages/pigeon
#import <Foundation/Foundation.h>
@protocol FlutterBinaryMessenger;
@class FlutterError;
@class FlutterStandardTypedData;

NS_ASSUME_NONNULL_BEGIN

@class FlutterResponse;
@class FlutterRequest;

@interface FlutterResponse : NSObject
@property(nonatomic, copy, nullable) NSString * message;
@end

@interface FlutterRequest : NSObject
@property(nonatomic, copy, nullable) NSString * expression;
@end

@interface FlutterCallbackApi : NSObject
- (instancetype)initWithBinaryMessenger:(id<FlutterBinaryMessenger>)binaryMessenger;
- (void)apply:(FlutterResponse*)input completion:(void(^)(NSError* _Nullable))completion;
@end
@protocol FlutterApi
-(nullable FlutterResponse *)call:(FlutterError *_Nullable *_Nonnull)error;
-(nullable FlutterResponse *)eval:(FlutterRequest*)input error:(FlutterError *_Nullable *_Nonnull)error;
@end

extern void FlutterApiSetup(id<FlutterBinaryMessenger> binaryMessenger, id<FlutterApi> _Nullable api);

NS_ASSUME_NONNULL_END