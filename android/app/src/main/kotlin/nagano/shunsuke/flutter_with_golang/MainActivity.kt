package nagano.shunsuke.flutter_with_golang

import androidx.annotation.NonNull
import io.flutter.embedding.android.FlutterActivity
import io.flutter.embedding.engine.FlutterEngine
import nagano.shunsuke.plugins.Pigeon


class MainActivity: FlutterActivity() {
    override fun configureFlutterEngine(@NonNull flutterEngine: FlutterEngine) {
        super.configureFlutterEngine(flutterEngine)
        Pigeon.Api.setup(flutterEngine.dartExecutor.binaryMessenger,  Api())
    }
}
